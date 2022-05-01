import { ApolloProvider } from "@apollo/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { client } from "../graphql/client";
import { TeamCreateRoute } from "./routes/TeamCreateRoute";
import { TeamRoute } from "./routes/TeamRoute";

const NotFound: React.FC = () => {
  return <p>404 Not Found</p>;
};

const Routing: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="team/create" element={<TeamCreateRoute />} />
        <Route path="team/:id" element={<TeamRoute />} />
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
