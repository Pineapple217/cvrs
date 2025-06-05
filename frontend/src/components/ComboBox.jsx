import { h } from "preact";
import { useSignal } from "@preact/signals";
import { useEffect, useRef } from "preact/hooks";

/**
 * @param {{ options: string[], placeholder?: string, onSelect?: (value: string) => void }} props
 */
export function Combobox({ options, placeholder = "Select...", onSelect }) {
  const inputRef = useRef(null);
  const dropdownVisible = useSignal(false);
  const inputValue = useSignal("");
  const filteredOptions = useSignal(options);

  const handleInput = (e) => {
    inputValue.value = e.target.value;
    dropdownVisible.value = true;
    filteredOptions.value = options.filter((opt) =>
      opt.toLowerCase().includes(e.target.value.toLowerCase())
    );
  };

  const handleSelect = (value) => {
    inputValue.value = value;
    dropdownVisible.value = false;
    if (onSelect) onSelect(value);
  };

  const handleBlur = () => {
    setTimeout(() => {
      dropdownVisible.value = false;
    }, 100); // Delay to allow click on item
  };

  return (
    <div class="combobox">
      <input
        ref={inputRef}
        class="combobox-input"
        type="text"
        placeholder={placeholder}
        value={inputValue.value}
        onInput={handleInput}
        onFocus={() => (dropdownVisible.value = true)}
        onBlur={handleBlur}
      />
      {dropdownVisible.value && filteredOptions.value.length > 0 && (
        <ul class="combobox-dropdown">
          {filteredOptions.value.map((opt) => (
            <li class="combobox-item" onMouseDown={() => handleSelect(opt)}>
              {opt}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
