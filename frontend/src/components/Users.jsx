import { useFetchUsers } from "../lib/users";

export function Users() {
  const { users, loading: userLoading, error } = useFetchUsers();

  if (error.value) {
    return <div class="alert alert-danger">Error: {error.value.message}</div>;
  }

  return (
    <>
      <span>
        <h2>
          Users
          {userLoading.value && <div class="loader"></div>}
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
          {users.value.map((u) => (
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
