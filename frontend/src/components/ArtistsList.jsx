import { useInfiniteQuery } from "@tanstack/react-query";
import { useAuth } from "./AuthProvider";
import { getArtists } from "../lib/artists";
import { useEffect, useRef } from "preact/hooks";

const FETCH_COUNT = 1;

export function ArtistsList() {
  const { token } = useAuth();
  const {
    data,
    error,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isFetchingNextPage,
    status,
  } = useInfiniteQuery({
    queryKey: ["artists"],
    queryFn: async ({ pageParam }) =>
      getArtists(token, FETCH_COUNT * pageParam, FETCH_COUNT),
    initialPageParam: 0,
    getNextPageParam: (lastPage, pages, lastPageParam) => {
      if (lastPage.length === 0) return null;
      return lastPageParam + 1;
    },
  });

  const loadMoreRef = useRef(null);
  useEffect(() => {
    if (!hasNextPage || isFetchingNextPage) return;

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting) {
          console.log("next page trig");
          fetchNextPage();
        }
      },
      { threshold: 0.1 }
    );

    if (loadMoreRef.current) {
      observer.observe(loadMoreRef.current);
    }

    return () => {
      if (loadMoreRef.current) {
        observer.unobserve(loadMoreRef.current);
      }
    };
  }, [hasNextPage, isFetchingNextPage, fetchNextPage]);

  return status === "pending" ? (
    <div class="alert">Loading</div>
  ) : status === "error" ? (
    <div class="alert alert-danger">Error: {error.message}</div>
  ) : (
    <>
      {data.pages.map((artists, i) => (
        <>
          {artists.map((artist) => (
            <div>
              <p key={artist.id}>{artist.name}</p>
              <img
                src={
                  __BACKEND_URL__ +
                  "/i/" +
                  artist.edges.image.edges.proccesed_image.find(
                    (a) => a.dimentions === 265
                  ).id
                }
                alt=""
              />
            </div>
          ))}
        </>
      ))}
      <div ref={loadMoreRef} style={{ height: 1 }}>
        {isFetchingNextPage && <p>Loading more...</p>}
        {!hasNextPage && <p>No more artists</p>}
      </div>
    </>
  );
}
