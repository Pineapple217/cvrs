import { signal } from "@preact/signals";
import { useAuth } from "../components/AuthProvider";
import { useEffect } from "preact/hooks";

/**
 * @typedef {Object} User
 * @property {string} id
 * @property {string} username
 * @property {boolean} is_admin
 * @property {Date} created_at
 */

/** @type {import('@preact/signals').Signal<User[]>} */
const users = signal([]);

/** @type {import('@preact/signals').Signal<boolean>} */
const loading = signal(false);

/** @type {import('@preact/signals').Signal<Error|null>} */
const error = signal(null);

export function useFetchUsers() {
  const { token } = useAuth();

  useEffect(() => {
    if (!token) return;

    const fetchUsers = async () => {
      loading.value = true;
      error.value = null;

      try {
        const res = await fetch(__BACKEND_URL__ + "/auth/users", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        if (!res.ok) throw new Error(`API returned ${res.status}`);
        // users.value = await res.json();
        users.value = (await res.json()).map((user) => ({
          ...user,
          created_at: new Date(user.created_at),
        }));
      } catch (err) {
        error.value = err;
      } finally {
        loading.value = false;
      }
    };

    fetchUsers();
  }, [token]);

  return { users, loading, error };
}
