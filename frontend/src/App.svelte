<script>
  import { onMount } from "svelte";

  const fallbackStudent = {
    name: "Chinmoy",
    roll_number: "2023bcs52"
  };

  let student = fallbackStudent;
  let status = "Fetching backend data...";

  onMount(async () => {
    const hostname = window.location.hostname || "3.110.215.199";
    const apiUrl = `http://${hostname}:8080/student-details`;

    try {
      const response = await fetch(apiUrl);
      if (!response.ok) {
        throw new Error(`HTTP ${response.status}`);
      }

      student = await response.json();
      status = `Backend connected: ${apiUrl}`;
    } catch (error) {
      student = fallbackStudent;
      status = `Backend unavailable at ${apiUrl}`;
    }
  });
</script>

<svelte:head>
  <title>Student Details</title>
</svelte:head>

<main class="card">
  <p class="eyebrow">Full Stack Assignment</p>
  <h1>Student Information</h1>

  <div class="row">
    <span class="label">Name</span>
    <span class="value">{student.name}</span>
  </div>

  <div class="row">
    <span class="label">Roll Number</span>
    <span class="value">{student.roll_number}</span>
  </div>

  <p class="status">{status}</p>
</main>

<style>
  :global(body) {
    margin: 0;
    min-height: 100vh;
    display: grid;
    place-items: center;
    font-family: Georgia, "Times New Roman", serif;
    background:
      radial-gradient(circle at top left, rgba(180, 83, 9, 0.16), transparent 30%),
      linear-gradient(135deg, #f7f1e8 0%, #efe3cf 100%);
    color: #1f2933;
  }

  .card {
    width: min(92vw, 680px);
    padding: 32px;
    background: #fffdf8;
    border: 1px solid #e7dcc9;
    border-radius: 20px;
    box-shadow: 0 24px 60px rgba(86, 54, 20, 0.14);
  }

  .eyebrow {
    margin: 0 0 8px;
    color: #b45309;
    font-size: 0.95rem;
    letter-spacing: 0.08em;
    text-transform: uppercase;
  }

  h1 {
    margin: 0 0 24px;
    font-size: clamp(2rem, 5vw, 3.6rem);
    line-height: 1;
  }

  .row {
    display: flex;
    justify-content: space-between;
    gap: 16px;
    padding: 14px 0;
    border-top: 1px solid #e7dcc9;
  }

  .label {
    font-weight: 700;
  }

  .value {
    text-align: right;
  }

  .status {
    margin: 20px 0 0;
    font-size: 0.95rem;
    color: #6b7280;
  }
</style>
