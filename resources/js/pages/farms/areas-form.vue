<template lang="pug">
  .areas-form
    .modal-header
      span.h4.font-bold(v-if="area.uid")
        translate Update Area
      span.h4.font-bold(v-else)
        translate Add New Area
    .modal-body
      p.text-muted
        translate Area is a space where you grow your plants. It could be a seeding tray, a garden bed, or a pot or anything that describes the different physical locations in your facility.
      b-form(@submit.prevent="validateBeforeSubmit")
        .line.line-dashed.b-b.line-lg
        .form-group
          label#label-name(for="name")
            translate Area Name
          input#name.form-control(
            type="text"
            v-validate="'required|alpha_num_space|min:5|max:100'"
            :class="{'input': true, 'text-danger': errors.has('name') }"
            v-model="area.name"
            name="name"
          )
          span.help-block.text-danger(v-show="errors.has('name')") {{ errors.first('name') }}

        .form-group
          label#label-size
              translate Size

          .form-row
            .col-xs-4.col-sm-4.col-md-6
              input#size.form-control(
                type="text"
                v-validate="'required|decimal'"
                :class="{'input': true, 'text-danger': errors.has('size') }"
                v-model="area.size"
                name="size"
              )
              span.help-block.text-danger(v-show="errors.has('size')") {{ errors.first('size') }}
            .col-xs-8.col-sm-8.col-md-6
              select#size_unit.form-control(
                v-validate="'required'"
                :class="{'input': true, 'text-danger': errors.has('size_unit') }"
                v-model="area.size_unit"
                name="size_unit"
              )
                option(v-for="size_unit in options.size_units" :value="size_unit.key")
                  | {{ size_unit.label }}
              span.help-block.text-danger(v-show="errors.has('size_unit')")
                | {{ errors.first('size_unit') }}
        .form-row
          .col-xs-6
            .form-group
              label#label-type(for="type")
                translate Type
              select#type.form-control(
                v-validate="'required'"
                :class="{'input': true, 'text-danger': errors.has('type') }"
                v-model="area.type"
                name="type"
              )
                option(v-for="type in options.types" :value="type.key") {{ type.label }}
              span.help-block.text-danger(v-show="errors.has('type')") {{ errors.first('type') }}
          .col-xs-6
            .form-group
              label#label-location(for="location") Locations
              select#location.form-control(
                v-validate="'required'"
                :class="{'input': true, 'text-danger': errors.has('location') }"
                v-model="area.location"
                name="location"
              )
                option(v-for="location in options.locations" :value="location.key")
                  | {{ location.label }}
              span.help-block.text-danger(v-show="errors.has('location')")
                | {{ errors.first('location') }}
        .form-row
          .col-xs-6
            .form-group
              label#label-reservoir(for="reservoir") Select Reservoir
              select#reservoir.form-control(
                v-validate="'required'"
                :class="{'input': true, 'text-danger': errors.has('reservoir') }"
                v-model="area.reservoir_id"
                name="reservoir"
              )
                option(value = "") Please select reservoir
                option(v-for="reservoir in reservoirs" :value="reservoir.uid") {{ reservoir.name }}
              span.help-block.text-danger(v-show="errors.has('reservoir')")
                | {{ errors.first('reservoir') }}
          .col-xs-6
            .form-group
              label Select photo
                small.text-muted (if any)
              UploadComponent(@fileSelelected="fileSelelected")
        .form-group
          BtnCancel(v-on:click.native="$parent.$emit('close')")
          BtnSave(customClass="float-right")
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import { AreaTypes, AreaLocations, AreaSizeUnits } from '../../stores/helpers/farms/area';
import { StubArea, StubMessage } from '../../stores/stubs';
import UploadComponent from '../../components/upload.vue';
import BtnCancel from '../../components/common/btn-cancel.vue';
import BtnSave from '../../components/common/btn-save.vue';

export default {
  name: 'FarmAreasForm',
  components: {
    UploadComponent,
    BtnCancel,
    BtnSave,
  },
  props: ['data'],
  data() {
    return {
      message: Object.assign({}, StubMessage),
      area: Object.assign({}, StubArea),
      filename: 'Не выбрано файлов',
      options: {
        types: Array.from(AreaTypes),
        locations: Array.from(AreaLocations),
        size_units: Array.from(AreaSizeUnits),
      },
    };
  },
  computed: {
    ...mapGetters({
      reservoirs: 'getAllReservoirs',
    }),
  },
  mounted() {
    this.fetchReservoirs();
    if (typeof this.data.uid !== 'undefined') {
      this.getAreaByUid(this.data.uid)
        .then(({ data }) => {
          this.area = data;
          this.area.size_unit = data.size.unit.symbol;
          this.area.size = data.size.value;
          this.area.location = data.location.code;
          this.area.reservoir_id = data.reservoir.uid;
        })
        .catch(error => error);
    } else {
      this.area.size_unit = this.options.size_units[0].key;
      this.area.location = this.options.locations[0].key;
      this.area.type = this.options.types[0].key;
    }
  },
  methods: {
    ...mapActions([
      'submitArea',
      'fetchReservoirs',
      'getAreaByUid',
    ]),
    validateBeforeSubmit() {
      this.$validator.validateAll().then((result) => {
        if (result) {
          this.submit();
        }
      });
    },
    submit() {
      this.submitArea(this.area)
        .then(() => this.$parent.$emit('close'))
        .catch(() => this.$toasted.error('Error in area submission'));
    },
    fileSelelected(file) {
      this.area.photo = file;
    },
  },
};
</script>
