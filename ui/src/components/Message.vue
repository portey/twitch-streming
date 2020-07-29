<template>
  <div>
    <strong>id: {{ message.id }}, type: {{ message.type }}, from: {{ message.subscriberName }}</strong>
  </div>
</template>

<script>
export default {
  props: {
    message: {
      type: Object,
    },
  },
};
</script>
