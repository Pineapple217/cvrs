import { signal } from "@preact/signals";
import { Header } from "../../components/Header";
import { Combobox } from "../../components/ComboBox";
import "./style.css";

const options = ["Apple", "Banana", "Orange", "Pear", "Grape"];
export function Home() {
  const handleSelect = (val) => {
    console.log("Selected:", val);
  };
  return (
    <main>
      <Header />
      <input type="text" />
      <div style="padding: 20px">
        <h3>Choose a fruit:</h3>
        <Combobox
          options={["Apple", "Banana", "Orange", "Grape", "Mango"]}
          placeholder="Pick a fruit"
          onSelect={(value) => console.log("Selected:", value)}
        />
        <input type="text" />
      </div>
    </main>
  );
}
