import { useSearchParams } from "react-router-dom";

export default function Cancel() {
  const [params] = useSearchParams();
  const sessionToken = params.get("token") ?? "";
  const cancelSubscription = async () => {
    const validSessionToken = await fetch(
      `${import.meta.env.VITE_BACKEND_ADDRESS}/auth/session/verify`,
      {
        method: "GET",
        mode: "cors",
        headers: {
          Authorization: sessionToken,
        },
      }
    )
      .then((resp) => {
        return resp.ok;
      })
      .catch(() => false);

    if (validSessionToken) {
      await fetch(
        `${import.meta.env.VITE_BACKEND_ADDRESS}/subscription/cancel`,
        {
          method: "POST",
          mode: "cors",
          headers: {
            Authorization: sessionToken,
          },
        }
      )
        .then((resp) => resp.json())
        .then((resp) => {
          console.log(resp);
        })
        .catch((err) => {
          console.error(err);
        });
    }
  };
  return (
    <>
      <button onClick={cancelSubscription}>Cancel Subscription</button>
    </>
  );
}
