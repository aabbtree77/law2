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

- No templates, no bs (Messenger chat popups, extra scroll bars, attention hijacking).

## Tech/Setup

1. Install node via nvm or volta.

2. Install [Tailwind CSS](https://tailwindcss.com/docs/installation).

3. Install [Tailwind CSS Typography plugin](https://tailwindcss.com/docs/typography-plugin).

4. Run

    ```
    ./gocode/md2html
    npx tailwindcss -i index-in.css -o index.css --watch
    cp ./lang/en/index.html . && gedit index.html
    ```   

The English language file is chosen as the main one, edit index.css and simonas.ico paths in the copied top level variant. Replace the section
titled by the HTML comment "Language radio" with the content of radio_toplevel.txt.

## Philosophy/Experience

The first rule of frontend is: Don't do the frontend.

[List is all you need](https://dynomight.net/lists/).

Instagram and X can be more convenient and direct.

Web design evolves like those colorful peafowls. This is about Darwin, not SOLID. I see people still add carousels with videos and "contact us" cards and chat assistants, progressive loading, all this extra clicking and attention hijacking. Horizontal scrolling just because "we can". Just drop it. None of this improves search, loading, information gathering. 

Leonardo da Vinci vs. [leonardo.ai](https://leonardo.ai/).

## Credits

Thanks to Samuel Dawson for his minimal [typewriter effect](https://tailwindflex.com/@samuel33/typewriter-effect) in vanilla Js. One must wrap the text div inside another div of a fixed height to avoid the side effect of the outer container resizing. 

DeepL is incredible, but it does not support Latin, Esperanto, or Japanese yet. However, it does support the Greek language.

I have applied two icons from [Heroicons](https://heroicons.com/). It is very nice to copy/paste svgs directly, too bad they have only a few icons. I am not sure icons are needed. Canon vs. Sony so to speak.

## P.S.

HTML is horrid. One can already write components with, say, [React (jsx)](https://www.youtube.com/watch?v=g9hPa-G3lfw), but that adds hundreds of kilobytes of needless Js to load and does not solve the layout problem. What helps a lot is simply fold/unfold divs in vscode with ctrl+shift+[]. ctrl+s not only saves a file, it formats a mess with prettier. Still no clue where that closing div is missing...

I did not have enough energy-time to try web components yet. The best components are nonexisting components.

Responsive layouts remain a huge pain. So many microproblems, so much micromanagement. F12 and resizing the window is not enough to reveal the bugs that occur on some real screens. There was one real screen overlapped text problem I could only simulate by changing the desktop resolution to 1280x720 and hitting F11.

There must be a better way to do HTML layouts, we need an open source engine for that, some kind of Blender HTML. Until those better times, I would just use single-column layouts, mobile first and mostly. Do not use h-screen at all, do not fit stuff into a desktop screen as they are all different, use very few font sizes, do not add any "hero cards" that use text with images nearby. Do not focus on the looks, get an MVP and iterate a bit, it is all about exponentially diminishing returns afterwards.
