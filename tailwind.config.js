module.exports = {
    content: [
        './templates/*.tmpl'
    ],
    theme: {
        extend: {
            keyframes: {
                'blink': {
                    '0%': { opacity: 1 },
                    '50%': { opacity: 0 },
                    '100%': { opacity: 1 }
                }
            },
            animation: {
                'header-arrow': 'blink 1.2s infinite linear'
            }
        }
    }
}