## Introduction 

This is a rewrite of [lawlt.eu](http://www.lawlt.eu/), a website of a law firm.

It solves the following problems:

- Uses HTTPS instead of HTTP.

- Removes the need for paid hosting, i.e. dropping a CMS and using public github pages as all the information is static and public anyway.

- More legible typography.

- Fixes a lot of i18n problems (9 languages, all practised by a single lawyer!)

## Setup

1. Install node via nvm or volta.

2. Install [Tailwind CSS](https://tailwindcss.com/docs/installation).

3. Install [Tailwind CSS Typography plugin](https://tailwindcss.com/docs/typography-plugin).

4. Install [daisyUI](https://daisyui.com/docs/install/).

5. Run

    ```
    ./gocode/md2html
    npx tailwindcss -i index-in.css -o index.css --watch
    ```

## Philosophy/Rant

The first rule of frontend is: Don't do the frontend.

[Lists is all you need](https://dynomight.net/lists/).

Web-wise, we are not viewers or readers, we are scanners.

Lately, I also notice some gyms inform people about their working schedule changes on the company's instagram account, not updating the homepage anymore. Instagram is more convenient, direct, and reaching more.

Web design evolves like those colorful peafowls. This is about Darwin, not SOLID. I see people still add carousels with videos and "contact us" cards and chat assistants, "progressive loading", and "cards" that demand extra clicking. We go for horizontal scrolling just because everyone is using vertical scrolling. None of this improves search, loading, information gathering. 

Leonardo da Vinci vs. [leonardo.ai](https://leonardo.ai/).
