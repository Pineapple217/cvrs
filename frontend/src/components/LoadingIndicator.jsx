import { useIsFetching } from "@tanstack/react-query";

export function LoadingIndicator() {
  const isFetching = useIsFetching();
  return (
    <div
      id="loading-line"
      style={{ display: isFetching ? "block" : "none" }}></div>
  );
}
