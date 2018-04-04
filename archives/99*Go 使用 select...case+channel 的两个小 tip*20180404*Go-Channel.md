这两个是我在面试中被问到的两个问题, 其中第二道我在面试时完全没想出, 之后知道答案之后觉得还是挺有意思的, 所以就在这儿记下来吧.

首先我们都知道在 Go 中使用 `select...case` 操作接收 channel 的时候是随机的, 假设现在有 `ch1` 和 `ch2` 两个channel, 那么现在下面这两个操作分别该怎么实现?

1. 实现 ch1 的优先级高于 ch2, 即两者同时可以被接收的时候, 绝对走 ch1 的 case.
2. 保持 select...case 的动态平衡, 即在两者同时被接收的时候, 最后选择两个 case 的概率大致相等.

---

首先我们来解决第一个问题.

这个问题还是比较简单的, 因为我们知道, 在进行 select...case 操作的时候, 虽然各个 case 的优先级一致, 但是这些操作的优先级都是比 default 要高的, 所以我们可以把对 ch2 的判断放到 default 里, 这样就实现了两者的优先级区别.

源码如下:

    select {
	case <-ch1:
		// do something
	default:
		select {
		case <-ch2:
			// do something else
		}
	}

---

我们再来看第二题, 这道题就有点意思了, 首先两个 case 不平衡的状况容易判断, 那么在不平衡发生的时候, 我们保证多的那一方被阻塞就好了, 所以这时候我们可以动态生成一个永远阻塞在接收状态的 channel, 什么样的 channel 符合这个条件呢?

1. nil channel;
2. 没有发送操作的 channel.

也就是说

> 直接接收会死锁 deadlock 的 channel 在 select...case 中会一直保持阻塞状态.

那么这样解法就好说了, 源码如下:

	getCh := func(i int, ch <-chan struct{}) <-chan struct{} {
		if i > 0 {
			return nil
		}
		return ch
	}

	i := 0
	for {
		select {
		case <-getCh(i, ch1):
			// do something
			i++
		case <-ch2:
			// do something else
			i--
		}
	}