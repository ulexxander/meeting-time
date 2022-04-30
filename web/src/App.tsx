import { ApolloProvider } from "@apollo/client";
import { client } from "../graphql/client";
import { useTeamByIdQuery } from "../graphql/generated";

const Team: React.FC<{ id: number }> = ({ id }) => {
  const { data, error } = useTeamByIdQuery({
    variables: {
      id: id.toString(),
    },
  });

  if (error) {
    return (
      <p>
        Error: {error.name} {error.message}
      </p>
    );
  }

  if (!data) {
    return <p>Loading...</p>;
  }

  if (!data.teamByID) {
    return <p>Team with id {id} does not exist</p>;
  }

  return (
    <p>
      Team {data.teamByID.id}: {data.teamByID.name}
    </p>
  );
};

const TeamsPage: React.FC = () => {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        flexDirection: "column",
        padding: 64,
        fontSize: 24,
      }}
    >
      <p style={{ textAlign: "center" }}>App!</p>
      <Team id={1} />
      <Team id={2} />
      <Team id={99} />
    </div>
  );
};

export const App: React.FC = () => {
  return (
    <ApolloProvider client={client}>
      <TeamsPage />
    </ApolloProvider>
  );
};
