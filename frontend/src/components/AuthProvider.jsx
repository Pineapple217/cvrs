import { createContext } from "preact";
import { useState, useEffect, useContext } from "preact/hooks";
import { useLocation } from "preact-iso";

/**
 * @typedef {Object} JwtPayload
 * @property {string} usn
 * @property {number} uid
 * @property {boolean} adm
 * @property {string} iss
 * @property {number} exp
 * @property {number} iat
 */

/**
 * @typedef {Object} AuthContextType
 * @property {string | null} token - Current JWT token
 * @property {(token: string) => void} setToken - Function to update the JWT token
 * @property {JwtPayload | null} payload - Decoded payload of the token
 */

/**
 * AuthContext provides authentication state and updater
 * @type {import('preact').Context<AuthContextType>}
 */
export const AuthContext = createContext({
  token: null,
  setToken: () => {},
  payload: null,
  loading: true,
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
  const [token, setTokenState] = useState(null);
  const [payload, setPayload] = useState(null);
  const [loading, setLoading] = useState(true);
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
    const stored = localStorage.getItem(__JWT_LOCALSTORAGE__);
    if (stored && isTokenValid(stored)) {
      setTokenState(stored);
      setPayload(decodePayload(stored));
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    if (!loading && !token && !url.startsWith("/auth/login")) {
      console.log(url);
      const redirectTo = encodeURIComponent(url);
      route(`/auth/login?r=${redirectTo}`);
    }
  }, [token, url, loading]);

  if (loading) {
    return null;
  }

  return (
    <AuthContext.Provider value={{ token, setToken, payload }}>
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
    const payload = decodePayload(token);
    return payload && Date.now() < payload.exp * 1000;
  } catch {
    return false;
  }
}

/**
 * Decode the payload of a JWT token
 * @param {string} token
 * @returns {JwtPayload}
 */
function decodePayload(token) {
  const [, payloadBase64] = token.split(".");
  const payloadJson = atob(payloadBase64);
  return JSON.parse(payloadJson);
}
