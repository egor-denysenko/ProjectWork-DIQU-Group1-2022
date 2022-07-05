"use strict";

const path = require("path");
const AutoLoad = require("@fastify/autoload");

module.exports = async function (fastify, opts) {
  // Place here your custom code!

  // Do not touch the following lines

<<<<<<< HEAD
	// This loads all plugins defined in plugins
	// those should be support plugins that are reused
	// through your application
	fastify.register(require('@fastify/cors'), {})
	fastify.register(AutoLoad, {
		dir: path.join(__dirname, "plugins"),
		options: Object.assign({}, opts),
	});
=======
  // This loads all plugins defined in plugins
  // those should be support plugins that are reused
  // through your application
  fastify.register(AutoLoad, {
    dir: path.join(__dirname, "plugins"),
    options: Object.assign({}, opts),
  });
>>>>>>> 1773b6ced36dae2e35cd88a27a7ba5425499ab42

  // This loads all plugins defined in routes
  // define your routes in one of these
  fastify.register(AutoLoad, {
    dir: path.join(__dirname, "routes"),
    options: Object.assign({}, opts),
  });
};
