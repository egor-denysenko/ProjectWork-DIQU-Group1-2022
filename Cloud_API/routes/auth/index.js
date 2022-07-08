"use strict";

const authData = require("./auth.data");
module.exports = async function (fastify, opts) {
  const psql = authData(fastify.sql);
  const saltRounds = 14;
  fastify.post("/", {
    schema: {
      tags: ["Login"],
      description:
        "Route For Verifying User To Login And Setting His User Id Into JWT",
      body: {
        type: "object",
        required: ["email", "password"],
        properties: {
          email: { type: "string" },
          password: { type: "string" },
        },
      },
    },
    response: {
      200: {
        type: "object",
        required: ["token", "username"],
        properties: {
          token: { type: "string" },
          username: { type: "string" },
        },
      },
    },
    handler: async function (request, reply) {
      try {
        const { email, password } = request.body;
        const login = await psql.login(email);
        const match = await fastify.bcrypt.compare(password, login["password"]);
        if (!match) {
          return [];
        } else {
          const jwtBearerToken = await fastify.jwt.sign(
            {
              id: login["id"],
              username: login["username"],
            },
            process.env.PRIVATE_RSA_KEY,
            {
              algorithm: "RS256",
              expiresIn: "6days",
            }
          );
          return { token: jwtBearerToken };
        }
      } catch (err) {
        console.log(err);
        console.log("errore e non manda nulla");
        return [];
      }
    },
  });
  fastify.post("/register", {
    schema: {
      tags: ["Login"],
      description: "Route For Creating A User",
      body: {
        type: "object",
        required: ["email", "username", "password"],
        properties: {
          email: { type: "string" },
          username: { type: "string" },
          password: { type: "string" },
        },
      },
    },
    response: {
      200: {
        type: "object",
        required: ["status"],
        properties: {
          status: { type: "string" },
        },
      },
    },
    handler: async function (request, reply) {
      try {
        const { email, username, password } = request.body;
        const hash = fastify.bcrypt.hashSync(password, saltRounds);
        console.log(psql);
        console.log("fastify psql robbe");
        await psql.signUp(email, username, hash);
        return { status: "Success" };
      } catch (err) {
        console.log(err);
        return { status: "Error On Creation" };
      }
    },
  }),
    fastify.post("/verify", {
      schema: {
        tags: ["Login"],
        description: "Verify If The JWT Is Correctly Signed",
        body: {
          type: "object",
          required: ["token"],
          properties: {
            token: { type: "string" },
          },
        },
      },
      response: {
        200: {
          type: "object",
          required: ["status"],
          properties: {
            status: { type: "boolean" },
          },
        },
      },
      handler: async function (request, reply) {
        try {
          const { token } = request.body;
          fastify.jwt.verify(token, process.env.PUBLIC_RSA_KEY, {
            algorithm: "RS256",
          });
          return { status: true };
        } catch (err) {
          return { status: false };
        }
      },
    });
};
