import { ApolloProvider } from "@apollo/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { client } from "../graphql/client";
import { TeamCreatePage } from "./pages/TeamCreatePage";
import { TeamPage } from "./pages/TeamPage";

const NotFound: React.FC = () => {
  return <p>404 Not Found</p>;
};

const Routing: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/team/create" element={<TeamCreatePage />} />
        <Route path="/team/:id" element={<TeamPage />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
};

export const App: React.FC = () => {
  return (
    <ApolloProvider client={client}>
      <Routing />
    </ApolloProvider>
  );
};
