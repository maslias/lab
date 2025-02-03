const { default: daisyui } = require("daisyui");

/** @type {import('tailwindcss').Config} */
module.exports = {
    // content: ["./views/*.html", "./views/*.templ", "./views/*.go"],
    content: ["./views/**/*.{html,templ,go}"],
    safelist: [],
    plugins: [require("@tailwindcss/typography"), require("daisyui")],
    theme: {
        fontFamily: {
            'sans': ['ui-sans-serif', 'system-ui'],
        }
    },
    daisyui: {
        themes: ["dark"], // false: only light + dark | true: all themes | array: specific themes like this ["light", "dark", "cupcake"]
        base: true, // applies background color and foreground color for root element by default
        styled: true, // include daisyUI colors and design decisions for all components
        utils: true, // adds responsive and modifier utility classes
        prefix: "", // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
        logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
        themeRoot: ":root", // The element that receives theme color CSS variables,
    },
};
