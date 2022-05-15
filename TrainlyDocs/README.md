# Website

This website is built using [Docusaurus 2](https://docusaurus.io/), a modern static website generator.

### Installation

```
$ pnpm install
```

### Local Development

```
$ pnpm start --port <PortNumber>
```

This command starts a local development server on the choosen port and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```
$ pnpm build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Deployment

Deployment To GitHub Pages

```
$ GIT_USER=<Your GitHub username> pnpm deploy
```

This command will only deploy the contend merged in the develop branch or in the main branch

So it's preferrable to finish writing the docs for one feature and after that merge and build the docs 

If you are using GitHub pages for hosting, this command is a convenient way to build the website and push to the `gh-pages` branch.
