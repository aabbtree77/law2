## Introduction 

This is a rewrite of [lawlt.eu](http://www.lawlt.eu/), a website of a law firm.

It solves the following problems:

- Adds a landing page. 

- Uses HTTPS instead of HTTP.

- Removes the need for a paid hosting, i.e. dropping a CMS and using public github pages as all the information is static and public anyway.

- A vastly better typography and color palette.

- Everything is now consistently in nine different languages, including the landing page. All the languages are practised by a single lawyer, by the way!

- Fixes a lot of i18n problems.

## Setup

1. Install node via nvm or volta.

2. Install [Tailwind CSS](https://tailwindcss.com/docs/installation).

3. Install [Tailwind CSS Typography plugin](https://tailwindcss.com/docs/typography-plugin).

## Technology/Philosophy

- Tailwind CSS: Extremely useful. Still somewhat low level, but makes everything local and you don't waste time navigating css files and naming classes.

- @tailwindcss/typography: A much better than default raw rendering of the results of the Markdown-to-HTML conversion. Not ideal, but saves a lot of time unless one is ready to manually set all these margins and paddings for h1 and p elements mingled with ul li in all sorts of combinatorial ways.

- daisyUI. I like its philosophy (everything should look good by default, focus on the looks and the end result, not whether you use @apply inside code). Planned to use a lot more of it, but applied only the radio button with a label in the end :). The card and avatar components could be useful, noted also the timeline component. Divider and other layout containers do not bring much as everything inside of them overextends by default unless properly adjusted. A very nice Jekyll-Hyde window split design somewhere in this lib, but still under a rewrite...  

- heroicons: A very nice idea to allow copy/pase svg directly, though HTML becomes horrid, and very few icons (around 400). Many slightly more abstract concepts like "newspaper", "wallet", "document", they do not look recognizable, or solving the problem of text replacement with images. 

Everything is still too low level in the year 2024, esp. it is so hard to adjust positions and sizes, div flex everywhere, width and height degrees of freedom unclear. h-1/3 may or may not work on a given element, no min-h-1/3, a global container size jumps once a tiny child element text content disappears and all sorts of low level crapola. Good to write a PhD, not so good to build a website under the clock ticking. 

## A Note on React

Tailwind and all these related libs are much better when used with React as we can abstract stuff into components and get a lot cleaner HTML.

Initially, experimented with vite, React, and shadcn which uses the embla carousel (now part of shadcn). Thought about going full scale with transparent layers, animations etc. Boy was I naive. Take a simple carousel. It is six layers of divs with plugins. Be my guest working with transparency, responsive layouts and accessibility. Very little documentation, and you start thinking about some extra tools to ease the debugging. 

Better do sophisticated custom landing in vanilla Js, or just ignore the bs. React is meaningful only when using the components en masse without having to style much.

Discarded React and shadcn as virtual DOM, jsx, and components turned to be totally unnecessary while adding some friction with configuration.

Also my frontend philosophy lately is to drop the nonsense as most of these elements do not improve search indexing, slow down rendering, and only hijack attention. There is really no need for horizontal scrolling, [lists is all you need](https://dynomight.net/lists/).

Concerning statics, react would bring more readable code, e.g.

```js
import { PhoneIcon } from "@heroicons/react/24/outline";
...
const Navbar = () => {
  return (
  ...
  <div className="flex flex-row items-center justify-around w-1/7">
          <PhoneIcon className="h-8 text-green-500" />
          <div className="font-bold text-2xl text-cyan-500">+370 650 74647</div>
        </div>
      </div>
    </nav>
  ...  
  );
};

export default Navbar;
```

Here the component PhoneIcon in jsx would be more readable than its copy-pasted svg block in html, but all this goodness comes at a price (jsx, config paths, @, loading times). 

Another react advantage is GUI (headless GUI) component libs, but one hardly needs any of that for static business landing pages. The only massively reused components would be "card" and "avatar" in shadcn, but they still need styling with all the pain of positioning and sizing around responsiveness.

I also noted the color scheme of react docs in the dark mode (for some future use), and this command (no need for vite when needing only tailwind):

npx tailwindcss -i ./index-in.css -o ./index.css --watch

## My Evolving Frontend Philosophy

The first rule of frontend is: Don't do the frontend.

Business landing pages are a contradiction. They are essentially just "yellow pages" with a documentation about the service, contact details, team etc. In reality, nobody wants this material to look like a documentation, so people waste time with a "unique" element placement and multiple responsive layered Figma designs, but most of this stuff looks good only big screens. It is not clear if layers and symmetry breaking add that much value, but they sure demand a lot of time and maybe even maintenance in the future. Ultimately, web-wise, we are not viewers or readers, we are scanners.

Lately, I also notice some gyms inform people about their working schedule changes on the company's instagram account, not updating the homepage anymore. Instagram is faster and reaching more.

Average business landing pages are so gross. I see people still add carousels with videos and "contact us" cards and chat assistants, "progressive loading" and iformation summary cards that demand extra clicking. We go for horizontal scrolling just because everyone is using vertical scrolling. None of this improves search, loading, information gathering. 

Web design evolves like those colorful peafowls. This is about Darwin, not SOLID.

Leonardo da Vinci vs. [leonardo.ai](https://leonardo.ai/).

In any case, I tried my best.

