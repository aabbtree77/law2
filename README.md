## Introduction 

<table>
<tr>
<th style="text-align:center">wepik AI generated cyberpunk lawyer based on the text prompt: Beauty wearing a judge hat, upscaled Jan 2024</th>
</tr>
<tr>
<td>
<img src="wepik-beauty-wearing-a-judge-hat-upscaled-2024Jan.jpg"  alt="wepik AI generated cyberpunk lawyer based on the text prompt: Beauty wearing a judge hat, upscaled 2024Jan" width="100%" >
</td>
</tr>
</table>

[Lawtrust.eu](https://lawtrust.eu/) is a rewrite of [lawlt.eu](http://www.lawlt.eu/), a law firm website.

My work solves the following problems:

- Uses HTTPS instead of HTTP.

- Removes the need for a paid hosting (dropping a CMS).

- Arguably better typography, mobile UX, i18n.

## Tech/Setup

1. Install node via nvm or volta.

2. Install [Tailwind CSS](https://tailwindcss.com/docs/installation).

3. Install [Tailwind CSS Typography plugin](https://tailwindcss.com/docs/typography-plugin).

4. Run

    ```
    cd gocode
    go build
    cd ..
    ./gocode/md2html
    npx tailwindcss -i index-in.css -o index.css --watch
    cp ./lang/en/index.html . && gedit index.html
    ```   

The English language file is chosen as the main one, edit `index.css`, `simonas.ico`, and `aealogo.jpg` paths in the copied top level `index.html`. Inject/copy-paste the content of radio_toplevel.txt below `<!-- Language radio -->` inside `index.html`.

## Take Away Messages

- A small non-marketed web site is almost invisible on the search. SEO is rain water. You are visible as long as you pay.

- A page on Instagram will bring higher visibility and will be easier to update.

- i18n is a maintenance burden. It looks nice and all, but 18 sections times 9 languages is tedious. Batch processing with DeepL is very limited.

- Modern web pages are full of nonsense: Carousels with videos, "Contact Us" cards, chat assistants, progressive loading, horizontal scrolling, modal windows...

- [List is all you need](https://dynomight.net/lists/). Find out how to make them a bit more fancy, that is it. Still need to be careful with a background image as it may add a white blink on the first paint while the page loads.

- Styling is hard as we do not have a fixed screen size and the same DPI. Redmi 12C (Android 14) has a browser viewport of 360x698 pixels and size 6.7". At the other end we have the UHD 3840×2160 resolution and size 32". Landscape vs Portrait...

- Relying on Chrome, F12, and window resizings may not be enough. I had an overlapping text which could only be simulated by changing the desktop resolution to 1280x720 and hitting F11. Resizing browser window after F12 would not reveal it.

- I would recommend mostly single-column layouts. Do not use h-screen at all, do not fit stuff into a desktop screen as they are all totally different. Use very few font sizes, do not add any hero cards, toasts, hamburger buttons, do not maximize the screen space usage. Do not focus on the looks, the CSS will be a micromanagement nightmare.

## Credits

Thanks to Samuel Dawson for his minimal [typewriter effect](https://tailwindflex.com/@samuel33/typewriter-effect) in vanilla Js. One must wrap the text div inside another div of a fixed height to avoid the side effect of the outer container resizing. 

I have applied two icons from [Heroicons](https://heroicons.com/). It is very nice to copy/paste svgs directly, too bad they have only a few icons. I am not sure icons are needed. Canon vs. Sony so to speak.

DeepL is incredible, but some checking and cleaning is generally needed, which can be annoying. Still, a decent translation. I think human translators will soon be needed only for high quality texts or special languages and dialects, e.g. Latin, Esperanto, [Samogitian](https://www.youtube.com/watch?v=lqSfOYhctoE)...




