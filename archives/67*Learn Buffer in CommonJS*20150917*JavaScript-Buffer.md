# Learn Buffer in CommonJS

A code snippet used to understand ```Buffer``` in CommonJS:

    var fs = require('fs');

    var rs = fs.createReadStream('test.cpp');
    var buffers = [];
    var nread = 0;

    rs.on('data', function (chunk) {
      // collect all buffer in the read stream to arr
      buffers.push(chunk);
      nread += chunk.length;
    });

    rs.on('end', function () {
      var buffer = null;
      switch (buffers.length) {
        case 0:
          buffer = new Buffer(0);
          break;
        case 1:
          buffer = buffers[0];
          break;
        default:
          buffer = new Buffer(nread);
          //  the code below will join the buffer into new buffer
          //  .copy method is to copy a buffer to a new buffer at target place
          //  and after join the buffers, we can use .toString to convert it to str, default character set is utf-8
          for (var i = 0, pos = 0, l = buffers.length; i < l; i ++) {
            var chunk = buffers[i];
            chunk.copy(buffer, pos);
            pos += chunk.length;
          }
          break;
      }
      console.log(buffer.toString());
    });
        