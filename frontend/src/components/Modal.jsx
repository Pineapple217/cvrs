import { effect } from "@preact/signals";

/**
 * A reusable modal component controlled via a signal.
 * @param {Object} props
 * @param {import('@preact/signals').Signal<boolean>} props.visible - A signal that controls modal visibility
 * @param {preact.ComponentChildren} props.children - The modal content
 * @param {() => void} [props.onClose] - Optional close handler
 */
export function Modal({ visible, children, onClose }) {
  effect(() => {
    if (!visible.value) return;

    /** @param {KeyboardEvent} e */
    const handler = (e) => {
      if (e.key === "Escape") {
        if (onClose) onClose();
        visible.value = false;
      }
    };

    document.addEventListener("keydown", handler);
    return () => document.removeEventListener("keydown", handler);
  });

  return (
    <div
      class="modal"
      style={{ display: visible.value ? "flex" : "none" }}
      onClick={() => {
        if (onClose) onClose();
        visible.value = false;
      }}>
      <div class="modal-content" onClick={(e) => e.stopPropagation()}>
        <div style={{ clear: "both" }}>{children}</div>
      </div>
    </div>
  );
}
