<template lang="pug">
  .row
    .col.text-truncate
      label.btn.btn-default.btn-file
        translate Browse
        input(type="file" @change="processFile($event)" style="display: none;")
      span.text-muted {{ filename }}
</template>

<script>
export default {
  name: 'UploadComponent',
  data() {
    return {
      filename: 'Не выбрано файлов',
    };
  },
  methods: {
    processFile(event) {
      this.filename = event.target.files[0].name;
      this.$emit('fileSelelected', event.target.files[0]);
    },
  },
};
</script>
