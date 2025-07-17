import { useIsFetching } from "@tanstack/react-query";

export function LoadingIndicator() {
  const isFetching = useIsFetching();
  return (
    <div id="loading-line" style={{ opacity: isFetching ? "1" : "0" }}></div>
  );
}
