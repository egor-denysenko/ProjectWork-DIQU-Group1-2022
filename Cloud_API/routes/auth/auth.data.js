const authDataAccess = (cockroach) => {
  const login = async (email) => {
    try {
      const user = await cockroach.one(
        'SELECT id,username,password FROM public."Users" WHERE email = $1',
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
      await cockroach.one(
        'INSERT INTO public."Users" (email, username, password) VALUES ($1,$2,$3) RETURNING id',
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
