const authDataAccess = (psql) => {
  const login = async (email) => {
    try {
      const user = await psql.one(
        "SELECT id,username,password FROM public.operators WHERE email = $1",
        [email]
      );
      return user;
    } catch (err) {
      console.log(err);
      throw err;
    }
  };
  const signUp = async (email, username, hash) => {
    try {
      await psql.one(
        "INSERT INTO public.operators (email, username, password) VALUES ($1,$2,$3) RETURNING id",
        [email, username, hash]
      );
    } catch (err) {
      console.log(err);
      throw err;
    }
  };

  return { login, signUp };
};

module.exports = authDataAccess;
