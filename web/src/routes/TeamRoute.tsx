import { useParams } from "react-router";
import { useTeamByIdQuery } from "../../graphql/generated";
import { Page } from "../components/layout";

const Team: React.FC<{ id: number }> = ({ id }) => {
  const { data, error } = useTeamByIdQuery({
    variables: {
      id,
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

const TeamByID: React.FC = () => {
  const { id = "" } = useParams<"id">();
  const idInt = parseInt(id);

  if (Number.isNaN(idInt)) {
    return <p>Invalid team id: {id}</p>;
  }

  return <Team id={idInt} />;
};

export const TeamRoute: React.FC = () => {
  return (
    <Page>
      <div className="text-xl">
        <p>This is the team</p>
        <TeamByID />
      </div>
    </Page>
  );
};
