// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require("prism-react-renderer/themes/github");
const darkCodeTheme = require("prism-react-renderer/themes/dracula");

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "Spellshape Cli Docs",
  tagline: "Spellshape Cli Docs",
  url: "https://docs.spellshape.com",
  baseUrl: "/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon-svg.svg",
  trailingSlash: false,

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "spellshape",
  projectName: "spellshape docs",

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  scripts: [
    {
      async: true,
      src: "https://www.googletagmanager.com/gtag/js?id=G-XL9GNV1KHW",
    },
  ],

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          versions: {
            current: {
              label: "nightly",
              path: "nightly",
              badge: true,
              banner: "unreleased", // put 'none' to remove
            },
          },
          sidebarPath: require.resolve("./sidebars.js"),
          routeBasePath: "/",
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      image: "img/og-image.jpg",
      announcementBar: {
        content:
          '<a target="_blank" rel="noopener noreferrer" href="https://spellshape.com">← Back to Spellshape</a>',
        isCloseable: false,
      },
      docs: {
        sidebar: {
          autoCollapseCategories: true,
        },
      },
      navbar: {
        hideOnScroll: true,
        title: "Spellshape",
        items: [
          {
            href: "https://github.com/spellshape/cli",
            html: `<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="github-icon">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M12 0.300049C5.4 0.300049 0 5.70005 0 12.3001C0 17.6001 3.4 22.1001 8.2 23.7001C8.8 23.8001 9 23.4001 9 23.1001C9 22.8001 9 22.1001 9 21.1001C5.7 21.8001 5 19.5001 5 19.5001C4.5 18.1001 3.7 17.7001 3.7 17.7001C2.5 17.0001 3.7 17.0001 3.7 17.0001C4.9 17.1001 5.5 18.2001 5.5 18.2001C6.6 20.0001 8.3 19.5001 9 19.2001C9.1 18.4001 9.4 17.9001 9.8 17.6001C7.1 17.3001 4.3 16.3001 4.3 11.7001C4.3 10.4001 4.8 9.30005 5.5 8.50005C5.5 8.10005 5 6.90005 5.7 5.30005C5.7 5.30005 6.7 5.00005 9 6.50005C10 6.20005 11 6.10005 12 6.10005C13 6.10005 14 6.20005 15 6.50005C17.3 4.90005 18.3 5.30005 18.3 5.30005C19 7.00005 18.5 8.20005 18.4 8.50005C19.2 9.30005 19.6 10.4001 19.6 11.7001C19.6 16.3001 16.8 17.3001 14.1 17.6001C14.5 18.0001 14.9 18.7001 14.9 19.8001C14.9 21.4001 14.9 22.7001 14.9 23.1001C14.9 23.4001 15.1 23.8001 15.7 23.7001C20.5 22.1001 23.9 17.6001 23.9 12.3001C24 5.70005 18.6 0.300049 12 0.300049Z" fill="currentColor"/>
            </svg>
            `,
            position: "right",
          },
          {
            href: "https://spellshape.com",
            className: "spellshape-backlink",
            label: `Back to Spellshape`,
            position: "right",
          },
          {
            type: "docsVersionDropdown",
            position: "left",
            dropdownActiveClassDisabled: true,
          },
        ],
      },
      footer: {
        links: [
          {
            items: [],
          },
          {
            title: "Products",
            items: [
              {
                label: "CLI",
                href: "https://spellshape.com/cli",
              },
              {
                label: "Accelerator",
                href: "https://spellshape.com/accelerator",
              },
            ],
          },
          {
            title: "Company",
            items: [
              {
                label: "About Spellshape",
                href: "https://spellshape.com/about",
              },
              {
                label: "Careers",
                href: "https://spellshape.com/careers",
              },
              {
                label: "Blog",
                href: "https://spellshape.com/blog",
              },
              {
                label: "Press",
                href: "https://spellshape.com/press",
              },
            ],
          },
          {
            title: "Contact",
            items: [
              {
                label: "Business Inquiries",
                href: "mailto:business@spellshape.com",
              },
            ],
          },
          {
            title: "Social",
            items: [
              {
                label: "Discord",
                href: "https://discord.com/invite/spellshape",
              },
              {
                label: "Twitter",
                href: "https://twitter.com/spellshape_com",
              },
              {
                label: "Linkedin",
                href: "https://www.linkedin.com/company/spellshape/",
              },
              {
                label: "YouTube",
                href: "https://www.youtube.com/spellshapehq",
              },
            ],
          },
        ],
        copyright: `<div>© Spellshape ${new Date().getFullYear()}</div><div><a href="https://spellshape.com/privacy">Privacy Policy</a><a href="https://spellshape.com/terms-of-use">Terms of Use</a></div>`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
        additionalLanguages: ["protobuf", "go-module"], // https://prismjs.com/#supported-languages
        magicComments: [
          // Remember to extend the default highlight class name as well!
          {
            className: "theme-code-block-highlighted-line",
            line: "highlight-next-line",
            block: { start: "highlight-start", end: "highlight-end" },
          },
          {
            className: "code-block-removed-line",
            line: "remove-next-line",
            block: { start: "remove-start", end: "remove-end" },
          },
        ],
      },
      algolia: {
        appId: "VVETP7QCVE",
        apiKey: "a9c466699c13052d35581030138f9fdc",
        indexName: "spellshape-cli",
        contextualSearch: false,
      },
      zoom: {
        selector: ".markdown :not(em) > img",
        config: {
          // options you can specify via https://github.com/francoischalifour/medium-zoom#usage
          background: {
            light: "rgb(255, 255, 255)",
            dark: "rgb(50, 50, 50)",
          },
        },
      },
    }),
  plugins: [
    [
      "@docusaurus/plugin-client-redirects",
      {
        redirects: [
          {
            from: "/nightly/guide/install",
            to: "/nightly/welcome/install",
          },
        ],
      },
    ],
    async function myPlugin(context, options) {
      return {
        name: "docusaurus-tailwindcss",
        configurePostCss(postcssOptions) {
          postcssOptions.plugins.push(require("postcss-import"));
          postcssOptions.plugins.push(require("tailwindcss/nesting"));
          postcssOptions.plugins.push(require("tailwindcss"));
          postcssOptions.plugins.push(require("autoprefixer"));
          return postcssOptions;
        },
      };
    },
    require.resolve("docusaurus-plugin-image-zoom"),
  ],
};

module.exports = config;
