// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Trainly',
  tagline: 'Modern solutions for train managment',
  url: 'https://egor-denysenko.github.io',
  baseUrl: '/ProjectWork-DIQU-Group1-2022/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  deploymentBranch:'gh-pages',
  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'egor-denysenko', // Usually your GitHub org/user name.
  projectName: 'ProjectWork-DIQU-Group1-2022', // Usually your repo name.

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'Trainly',
        logo: {
          alt: 'Trainly Logo',
          src: 'img/DocusaurusLogo.svg',
        },
        items: [
          {
            type: 'doc',
            docId: 'Introduction/Project_Brief',
            position: 'left',
            label: 'Docs',
          },
          {
            href: 'https://github.com/egor-denysenko/ProjectWork-DIQU-Group1-2022',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        copyright: `Copyright © ${new Date().getFullYear()} Trainly`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;
