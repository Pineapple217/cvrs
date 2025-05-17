import { Header } from "../../components/Header";
import { Users } from "../../components/Users";

export function Admin() {
  return (
    <main>
      <Header />
      <div class="container" style="max-width: 800px;">
        <Users />
      </div>
    </main>
  );
}
