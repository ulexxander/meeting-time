import { useParams } from "react-router";
import { TeamByIdQuery, useTeamByIdQuery } from "../../graphql/generated";
import { Page } from "../components/layout";

type TeamData = NonNullable<TeamByIdQuery["teamByID"]>;

const Schedules: React.FC<{ schedules: TeamData["schedules"] }> = ({
  schedules,
}) => {
  return (
    <ol className="list-decimal">
      {schedules.map((item) => (
        <li key={item.id}>
          <p>{item.name}</p>
          <p className="tracking-widest">
            {item.startsAt} - {item.endsAt}
          </p>
        </li>
      ))}
    </ol>
  );
};

const Team: React.FC<{ team: TeamData }> = ({ team }) => {
  const { id, name, schedules } = team;
  return (
    <div>
      <h3>{name}</h3>
      <p className="text-sm text-gray-600">Team #{id}</p>
      <div className="mt-4">
        <Schedules schedules={schedules} />
      </div>
    </div>
  );
};

const TeamLoader: React.FC<{ id: number }> = ({ id }) => {
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

  return <Team team={data.teamByID} />;
};

const TeamByID: React.FC = () => {
  const { id = "" } = useParams<"id">();
  const idInt = parseInt(id);

  if (Number.isNaN(idInt)) {
    return <p>Invalid team id: {id}</p>;
  }

  return <TeamLoader id={idInt} />;
};

export const TeamRoute: React.FC = () => {
  return (
    <Page>
      <div className="text-xl">
        <TeamByID />
      </div>
    </Page>
  );
};
