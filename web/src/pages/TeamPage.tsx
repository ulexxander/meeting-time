import { useParams } from "react-router";
import { useTeamByIdQuery } from "../../graphql/generated";

const Team: React.FC = () => {
  const { id = "" } = useParams<"id">();
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

export const TeamPage: React.FC = () => {
  return (
    <div className="flex flex-col items-center p-16 text-xl">
      <p>This is the team</p>
      <Team />
    </div>
  );
};
