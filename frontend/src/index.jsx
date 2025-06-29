import { render } from "preact";
import { LocationProvider, Router, Route } from "preact-iso";

import { Header } from "./components/Header.jsx";
import { AuthProvider } from "./components/AuthProvider.jsx";
import { Login } from "./components/Login.jsx";
import { Home } from "./pages/Home/index.jsx";
import { NotFound } from "./pages/_404.jsx";
import { Admin } from "./pages/Admin/index.jsx";
import "./style.css";
import { User } from "./pages/User/index.jsx";
import { Artists } from "./pages/Artists/index.jsx";

export function App() {
  return (
    <LocationProvider>
      <AuthProvider>
        <main>
          <body>
            <Router>
              <Route path="/auth/login" component={Login} />
              <Route path="/auth/user" component={User} />
              <Route path="/" component={Home} />
              <Route path="/artists" component={Artists} />
              <Route path="/admin" component={Admin} />
              <Route default component={NotFound} />
            </Router>
          </body>
        </main>
      </AuthProvider>
    </LocationProvider>
  );
}

render(<App />, document.getElementById("app"));
