import { createContext } from "preact";
import { useState, useEffect, useContext } from "preact/hooks";
import { useLocation } from "preact-iso";

/**
 * @typedef {Object} AuthContextType
 * @property {string | null} token - Current JWT token
 * @property {(token: string) => void} setToken - Function to update the JWT token
 */

/**
 * AuthContext provides authentication state and updater
 * @type {import('preact').Context<AuthContextType>}
 */
export const AuthContext = createContext({
  token: null,
  setToken: () => {},
});

/**
 * AuthProvider component wraps the app and manages auth state
 * It checks for a JWT in localStorage on mount, redirects to login if missing,
 * and stores an acquired token in sessionStorage and context.
 *
 * @param {{ children: import('preact').ComponentChildren }} props
 */
export function AuthProvider({ children }) {
  // Initialize token from localStorage
  const [token, setTokenState] = useState(() => {
    const stored = localStorage.getItem(__JWT_LOCALSTORAGE__);
    return stored && isTokenValid(stored) ? stored : null;
  });
  const { url, route } = useLocation();

  /**
   * Update token in state and sessionStorage
   * @param {string} newToken - JWT token string
   */
  const setToken = (newToken) => {
    localStorage.setItem(__JWT_LOCALSTORAGE__, newToken);
    setTokenState(newToken);
  };

  useEffect(() => {
    if (!token && url != "/auth/login") {
      route("/auth/login");
    }
  }, [token, url]);

  return (
    <AuthContext.Provider value={{ token, setToken }}>
      {children}
    </AuthContext.Provider>
  );
}

/**
 * Hook to access authentication context
 * @returns {AuthContextType}
 */
export function useAuth() {
  return useContext(AuthContext);
}

/**
 * Decode JWT and check if it's expired
 * @param {string} token
 * @returns {boolean}
 */
function isTokenValid(token) {
  try {
    const [, payloadBase64] = token.split(".");
    const payload = JSON.parse(atob(payloadBase64));
    const exp = payload.exp;
    if (!exp) return false;
    return Date.now() < exp * 1000; // exp is in seconds, JS uses ms
  } catch (e) {
    return false;
  }
}
