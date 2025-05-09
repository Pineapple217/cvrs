import { render } from "preact";
import { LocationProvider, Router, Route } from "preact-iso";

import { Header } from "./components/Header.jsx";
import { AuthProvider } from "./components/AuthProvider.jsx";
import { Login } from "./components/Login.jsx";
import { Home } from "./pages/Home/index.jsx";
import { NotFound } from "./pages/_404.jsx";
import "./style.css";

export function App() {
  return (
    <LocationProvider>
      <AuthProvider>
        {/* <Header /> */}
        <main>
          <Router>
            <Route path="/auth/login" component={Login} />
            <Route path="/" component={Home} />
            <Route default component={NotFound} />
          </Router>
        </main>
      </AuthProvider>
    </LocationProvider>
  );
}

render(<App />, document.getElementById("app"));
