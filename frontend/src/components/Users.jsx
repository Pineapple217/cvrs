import { useQuery } from "@tanstack/react-query";
import { getUsers } from "../lib/users";
import { useAuth } from "./AuthProvider";

export function Users() {
  const { token } = useAuth();
  const { error, data, isFetching } = useQuery({
    queryKey: ["users"],
    staleTime: 2 * 60 * 1000,
    queryFn: async () => getUsers(token),
  });

  if (error) {
    return <div class="alert alert-danger">Error: {error.message}</div>;
  }

  return (
    <>
      <span>
        <h2>
          Users
          {isFetching && <div class="loader"></div>}
        </h2>
      </span>
      <table>
        <thead>
          <tr>
            <th>Id</th>
            <th>Username</th>
            <th>Role</th>
            <th>Created At</th>
          </tr>
        </thead>
        <tbody>
          {data &&
            data.map((u) => (
              <tr key={u.id}>
                <td>{u.id}</td>
                <td>{u.username}</td>
                <td>{u.is_admin ? "Admin" : "User"}</td>
                <td>{u.created_at.toLocaleString()}</td>
              </tr>
            ))}
        </tbody>
      </table>
    </>
  );
}
