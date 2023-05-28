import { useSearchParams } from "react-router-dom";

export default function Cancel() {
  const [params] = useSearchParams();
  const sessionToken = params.get("token") ?? "";
  const cancelSubscription = async () => {
    const validSessionToken = await fetch(
      `${import.meta.env.VITE_BACKEND_ADDRESS}/auth/session/check`,
      {
        method: "GET",
        mode: "cors",
        headers: {
          Authorization: sessionToken,
        },
      }
    )
      .then(async (resp) => {
        const respBody = await resp.json();
        if (!resp.ok) {
          window.alert(respBody["errorMsg"]);
          return false;
        }
        return true;
      })
      .catch((err) => {
        window.alert(err);
        return false;
      });

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
        .then(async (resp) => {
          if (!resp.ok) {
            window.alert("Sad to see you go!");
            return;
          }
          window.open("https://checkout.apollo-music.app", "_self");
        })
        .catch((err) => {
          window.alert(err);
        });
    }
  };
  return (
    <>
      <button onClick={cancelSubscription}>Cancel Subscription</button>
    </>
  );
}
