import { useAuth } from "./AuthProvider";

export function Header() {
  const { payload } = useAuth();
  return (
    <header>
      <h1>
        <a href="/">CVRS</a>
      </h1>
      <nav style="flex: 1;">
        <ul>
          <li>
            <a href="/artists">Artists</a>
          </li>
          <li>
            <a href="/releases">Releases</a>
          </li>
          {payload?.adm && (
            <li>
              <a href="/admin">Admin</a>
            </li>
          )}
        </ul>
      </nav>
      <nav>
        <ul>
          <li>
            <a href={payload ? "/auth/user" : "/auth/login"}>
              {payload ? payload.usn : "Login"}
            </a>
          </li>
        </ul>
      </nav>
    </header>
  );
}
