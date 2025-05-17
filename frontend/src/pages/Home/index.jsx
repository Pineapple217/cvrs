import { useAuth } from "../../components/AuthProvider";
import { Header } from "../../components/Header";
import "./style.css";

export function Home() {
  const { payload } = useAuth();

  return (
    <main>
      <Header />
    </main>
  );
}
