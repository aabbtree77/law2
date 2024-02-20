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

This is a rewrite of [lawlt.eu](http://www.lawlt.eu/), a law firm website.

My work solves the following problems:

- Uses HTTPS instead of HTTP.

- Removes the need for paid hosting, i.e. dropping a CMS and using public github pages as all the information is static and public anyway.

- More legible typography.

- Better mobile UX (cough, cough, responsive layout).

- Fixes a lot of i18n problems (9 languages, all practised by a single lawyer!)

- Provides a unique quality landing page and promotes minimalism (no templates, no bs).

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

Web-wise, we are not viewers or readers, we are scanners.

Instagram and X are more convenient and direct.

Web design evolves like those colorful peafowls. This is about Darwin, not SOLID. I see people still add carousels with videos and "contact us" cards and chat assistants, progressive loading, all this extra clicking and attention hijacking. Horizontal scrolling just because "we can". Just drop it. None of this improves search, loading, information gathering. 

Leonardo da Vinci vs. [leonardo.ai](https://leonardo.ai/).

## Credits

Thanks to Samuel Dawson for his minimal [typewriter effect](https://tailwindflex.com/@samuel33/typewriter-effect) in vanilla Js. One must wrap the text div inside another div of a fixed height to avoid the side effect of the outer container resizing. 

DeepL is incredible, but it does not support Latin, Esperanto, or Japanese yet. However, it does support the Greek language.

I have applied two icons from [Heroicons](https://heroicons.com/). It is very nice to copy/paste svgs directly, too bad they have only a few icons. I am not sure icons are needed. Canon vs. Sony so to speak.

## P.S.

HTML is horrid, but one can already write components with, say, React. One can also simply fold/unfold divs in vscode with ctrl+shift+[]. ctrl+s not only saves a file, it formats a mess with prettier and lints the code. Often, one still does not know where that closing div is missing, but these two commands are already incredible and with vertical div matching lines they win over React components, in the static case.

I did not have enough energy-time to try web components yet. The best components are nonexisting components.

Styling remains a pain. So many microproblems, so much micromanagement. Very good for a PhD thesis, not so good for productivity. Hence, see the first rule of frontend above. 

I get an uneasy feeling that the Blender community achieves a lot more, but the grass is always greener on the other side. I suspect they still cannot drill a hole without manual retopology every time.

