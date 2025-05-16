import { useAuth } from "../../components/AuthProvider";
import "./style.css";

export function Home() {
  const { payload } = useAuth();
  console.log(payload?.usn); // "pine"

  return <div class="home">{payload && <h1>{payload.usn}</h1>}</div>;
}
