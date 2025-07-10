/**
 * @typedef {Object} User
 * @property {string} id
 * @property {string} username
 * @property {boolean} is_admin
 * @property {Date} created_at
 */

/**
 * @returns {Promise<User[]>}
 */
export const getUsers = async (token) => {
  const response = await fetch(__BACKEND_URL__ + "/auth/users", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  /** @type {Array<Omit<User, "created_at"> & { created_at: string }>} */
  const raw = await response.json();

  /** @type {User[]} */
  const users = raw.map((u) => ({
    ...u,
    created_at: new Date(u.created_at),
  }));

  return users;
};
