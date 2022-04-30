import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useTeamCreateMutation } from "../../graphql/generated";
import { Page } from "../components/layout";

const teamPageRedirectTimeout = 1000;

const TeamCreateForm: React.FC = () => {
  const { register, handleSubmit, formState } = useForm<{
    name: string;
  }>();
  const [teamCreate, { data, loading, error }] = useTeamCreateMutation();
  const navigate = useNavigate();

  useEffect(() => {
    if (data?.teamCreate) {
      const timeout = setTimeout(() => {
        navigate(`/team/${data.teamCreate}`);
      }, teamPageRedirectTimeout);
      return () => clearTimeout(timeout);
    }
    return () => {};
  }, [data]);

  return (
    <div>
      <form
        onSubmit={handleSubmit((fields) => {
          teamCreate({
            variables: {
              input: fields,
            },
          });
        })}
      >
        <div className="mb-6">
          <label
            htmlFor="team-create-name"
            className="mb-2 block text-sm font-medium text-gray-900"
          >
            Team name
          </label>
          <input
            type="text"
            id="team-create-name"
            className="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 outline-none focus:border-blue-500 focus:ring-blue-500"
            placeholder="Awesome team!"
            required
            {...register("name", { required: true })}
          />
          {formState.errors.name && (
            <span className="text-red-500">
              {formState.errors.name.message}
            </span>
          )}
        </div>

        <button
          type="submit"
          className="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 sm:w-auto"
        >
          Create team
        </button>
      </form>

      {loading && <p>Loading...</p>}

      {error && (
        <p className="text-red-500">
          Error creating team: {error.name} {error.message}
        </p>
      )}

      {data && <p className="text-green-500">Successfully created team!</p>}
    </div>
  );
};

export const TeamCreatePage: React.FC = () => {
  return (
    <Page>
      <h2>Create new team</h2>
      <div className="mt-6">
        <TeamCreateForm />
      </div>
    </Page>
  );
};
