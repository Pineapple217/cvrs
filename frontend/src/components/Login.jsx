import { useState } from "preact/hooks";
import { useAuth } from "./AuthProvider";

export function Login() {
  const { setToken } = useAuth();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);

  /**
   * Handle form submission
   * @param {import('preact').JSX.TargetedEvent<HTMLFormElement, Event>} e
   */
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    try {
      // Replace with your real auth endpoint
      const response = await fetch(__BACKEND_URL__ + "/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });
      if (!response.ok) throw new Error("Invalid credentials");
      const { token } = await response.json();
      // Save token via context (also sessionStorage inside)
      setToken(token);
      // Redirect to home or protected route
      window.location.replace("/");
    } catch (err) {
      setError(err.message);
    }
  };

  /**
   * Handle username input change
   * @param {import('preact').JSX.TargetedEvent<HTMLInputElement, Event>} e
   */
  const handleUsernameInput = (e) => {
    setUsername(e.currentTarget.value);
  };

  /**
   * Handle password input change
   * @param {import('preact').JSX.TargetedEvent<HTMLInputElement, Event>} e
   */
  const handlePasswordInput = (e) => {
    setPassword(e.currentTarget.value);
  };

  return (
    <div class="login-container">
      <form class="login-form" onSubmit={handleSubmit}>
        <h2>Login</h2>
        {error && <p class="error">{error}</p>}

        <label htmlFor="username">Username</label>
        <input
          id="username"
          name="username"
          type="text"
          value={username}
          onInput={handleUsernameInput}
          required
        />

        <label htmlFor="password">Password</label>
        <input
          id="password"
          name="password"
          type="password"
          value={password}
          onInput={handlePasswordInput}
          required
        />

        <button type="submit">Log In</button>
      </form>
    </div>
  );
}
