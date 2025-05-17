import { useAuth } from "../../components/AuthProvider";
import { Header } from "../../components/Header";

export function User() {
  const { payload } = useAuth();
  return (
    <main>
      <Header />
      <div class="container" style="max-width: 800px;">
        <h1>{payload.usn}</h1>
        <p>{payload.adm ? "Admin" : "User"}</p>
        <p>Id: {payload.uid}</p>
      </div>
    </main>
  );
}
