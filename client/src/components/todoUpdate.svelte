<script>
  let updateDescription = "";
  let updateTitle = "";

  //accepting props
  export let todo;
  export let editTodo;
  export let setEditButtonClick;

  const editCurrentTodo = () => {
    if (updateDescription != "" && updateTitle != "") {
      todo.description = updateDescription;
      todo.title = updateTitle;
      setEditButtonClick();
      editTodo(todo);
    }
  };
</script>

<div>
  <div class="flex flex-col sm:flex-row gap-2 m-4">
    <input
      bind:value={updateTitle}
      class="shadow border rounded w-full py-2 px-3 mr-4 text-gray-600"
      placeholder="New Title"
    />
    <input
      bind:value={updateDescription}
      class="shadow border rounded w-full py-2 px-3 mr-4 text-gray-600"
      placeholder="New Description"
    />
    <button
      on:click|preventDefault={editCurrentTodo}
      class="flex-no-shrink p-2 rounded-lg text-white bg-gray-500 hover:text-red hover:bg-gray-700"
      >Update</button
    >
  </div>
</div>
