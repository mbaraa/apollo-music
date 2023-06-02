import Signup from "./components/Signup";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Landing from "./components/Landing";
import Cancel from "./components/Cancel";
import Intro from "./components/Intro";

async function checkSession() {
  return await fetch(
    `${import.meta.env.VITE_BACKEND_ADDRESS}/auth/session/check`,
    {
      method: "GET",
      headers: {
        Authorization: localStorage.getItem("token") ?? "",
      },
    }
  )
    .then(async (resp) => {
      localStorage.setItem("token", (await resp.json())["data"]["token"]);
      return resp.ok;
    })
    .catch(() => false);
}

async function gotoApp() {
  if (await checkSession()) {
    window.open(
      `${import.meta.env.VITE_APP_ADDRESS}/sign-in?token=${localStorage.getItem(
        "token"
      )}`,
      "_self"
    );
  }
}

function App() {
  const router = createBrowserRouter([
    { path: "/", element: <Landing /> },
    { path: "/intro", element: <Intro /> },
    { path: "/sign-up", element: <Signup /> },
    { path: "/cancel", element: <Cancel /> },
  ]);

  gotoApp();

  return (
    <div className="font-[Comfortaa]">
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
