import { render } from "preact";
import { LocationProvider, Router, Route } from "preact-iso";

import { AuthProvider } from "./components/AuthProvider.jsx";
import { Login } from "./components/Login.jsx";
import { Home } from "./pages/Home/index.jsx";
import { NotFound } from "./pages/_404.jsx";
import { Admin } from "./pages/Admin/index.jsx";
import "./style.css";
import { User } from "./pages/User/index.jsx";
import { Artists } from "./pages/Artists/index.jsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Artist } from "./pages/Artist/index.jsx";
import { LoadingIndicator } from "./components/LoadingIndicator.jsx";
import { Releases } from "./pages/Releases/index.jsx";

const queryClient = new QueryClient();

export function App() {
  return (
    <LocationProvider>
      <AuthProvider>
        <QueryClientProvider client={queryClient}>
          <main>
            <LoadingIndicator />
            <body>
              <Router>
                <Route path="/auth/login" component={Login} />
                <Route path="/auth/user" component={User} />
                <Route path="/" component={Home} />
                <Route path="/artists" component={Artists} />
                <Route path="/artist/:id" component={Artist} />
                <Route path="/admin" component={Admin} />
                <Route path="/releases" component={Releases} />
                <Route default component={NotFound} />
              </Router>
            </body>
          </main>
        </QueryClientProvider>
      </AuthProvider>
    </LocationProvider>
  );
}

render(<App />, document.getElementById("app"));
