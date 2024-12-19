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

- Removes the need for a paid hosting. Dropping a CMS and using public github pages as all the information is static and public anyway.

- More legible typography.

- Better mobile UX (cough, cough, responsive layout).

- Fixes a lot of i18n problems (9 languages, all practised by a single lawyer!)

- No templates, no bs (chat popups, extra scroll bars, attention hijacking).

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

## Philosophy/Experience

The first rule of frontend is: Don't do the frontend.

[List is all you need](https://dynomight.net/lists/). The problem is that lists are also boring. Also, if there is a background image, one may encounter a white blink on the first paint while the page loads, despite the simplicity. 

A page on Instagram will be more convenient and direct than all this creativity. Think of information updates. I did this site myself on February, then I came back to update it on December and I would barely know what to do if I did not document most of my actions properly.
Already some new warnings to fix for Tailwind CSS. Lately I also use Bun more than Node. This Js ecosystem is just unbelievable. One year, two max, and everything rottens.

Also, a small non-marketed web site is almost invisible. SEO is rain water. You are visible as long as you pay Google/Microsoft. If you paid them in the past, you might still end up somewhere. If not, you might not even be searchable!

Web design can be horrid. I see people still add carousels with videos and "contact us" cards and chat assistants, progressive loading, all this extra clicking and attention hijacking. Horizontal scrolling and swipes just because "we can". Just drop it. None of this improves search, loading, information gathering.

## Credits

Thanks to Samuel Dawson for his minimal [typewriter effect](https://tailwindflex.com/@samuel33/typewriter-effect) in vanilla Js. One must wrap the text div inside another div of a fixed height to avoid the side effect of the outer container resizing. 

DeepL is incredible, but it does not support Latin and Esperanto. [Samogitian](https://www.youtube.com/watch?v=lqSfOYhctoE) will probably be out of the reach of these AI tools for a long long time...

I have applied two icons from [Heroicons](https://heroicons.com/). It is very nice to copy/paste svgs directly, too bad they have only a few icons. I am not sure icons are needed. Canon vs. Sony so to speak.

## P.S.

HTML/CSS are horrid. One can already write components with, say, [React (JSX)](https://www.youtube.com/watch?v=g9hPa-G3lfw), but that adds hundreds of kilobytes of needless Js to load and does not solve the layout problem. What helps a lot is simply fold/unfold divs in vscode with ctrl+shift+[]. ctrl+s not only saves a file, it formats a mess with prettier. Still no clue where that closing div is missing...

I did not have enough energy-time to try the web components yet. The best components are nonexisting components. Nobody knows how to define good boundaries for them, we are all just beta testers of these n+1 things. "We got it right this time..."

Responsive layouts remain a huge pain. So many microproblems, so much micromanagement. F12 and resizing the window is not enough to reveal the bugs that occur on some real screens. I had an overlapping text problem I could only simulate by changing the desktop resolution to 1280x720 and hitting F11. Resizing browser window after F12 would not reveal it.

There must be a better way to do HTML layouts, we need some kind of Blender HTML. Until those better times, I would just use single-column layouts, mobile first and mostly. Do not use h-screen at all, do not fit stuff into a desktop screen as they are all different, use very few font sizes, do not add any "hero cards" that use text with images nearby. Do not focus on the looks, get an MVP and iterate a bit, this is all about exponentially diminishing returns.
