import Grid from "@mui/material/Grid";
import { useMaterialUIController } from "../../../../context";
import JobCard from "../JobCard";

function JobGrid() {
  const [controller] = useMaterialUIController();
  const { jobs } = controller;
  console.log("#JobGrid: jobs", jobs);

  return (
    <Grid container spacing={2}>
      {jobs.map((job) => (
        <Grid key={`job_${job.id}`} item xs={12} md={6} lg={6} alignItems="stretch">
          <JobCard job={job} />
        </Grid>
      ))}
    </Grid>
  );
}

export default JobGrid;
