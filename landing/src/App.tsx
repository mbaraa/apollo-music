import Signup from "./components/Signup";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Landing from "./components/Landing";
import Cancel from "./components/Cancel";
import Intro from "./components/Intro";

function App() {
  const router = createBrowserRouter([
    { path: "/", element: <Landing /> },
    { path: "/intro", element: <Intro /> },
    { path: "/sign-up", element: <Signup /> },
    { path: "/cancel", element: <Cancel /> },
  ]);
  return (
    <div className="font-[Comfortaa]">
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
