export const InventoryTypes = [
  { key: 'seed',  label: 'Семя' },
  { key: 'growing_medium', label: 'Саженец' },
  { key: 'agrochemical', label: 'Агрохимия' },
  { key: 'label_and_crop_support', label: 'Label and Crop Support' },
  { key: 'seeding_container', label: 'Контейнер для проращивания' },
  { key: 'post_harvest_supply', label: 'Отправка после сбора урожая' },
  { key: 'plant', label: 'Растение' },
  { key: 'other', label: 'Другие материалы' },
]

export function FindInventoryType(key) {
  var inventoryType = InventoryTypes.find(item => item.key === key.toLowerCase())
  return inventoryType ? inventoryType.label : ''
}

export const QuantityUnits = [
  { key: 'SEEDS',  label: 'Семена' },
  { key: 'PACKETS', label: 'Пакеты' },
  { key: 'GRAM', label: 'Грамм' },
  { key: 'KILOGRAM', label: 'Килограмм' },
]

export function FindQuantityUnit(key) {
  var quantityUnit =  QuantityUnits.find(item => item.key === key.toLowerCase())
  return quantityUnit ? quantityUnit.label : ''
}

export const AgrochemicalQuantityUnits = [
  { key: 'PACKETS',  label: 'Пакеты' },
  { key: 'BOTTLES', label: 'Бутылки' },
  { key: 'BAGS', label: 'Мешки' },
]

export function FindAgrochemicalQuantityUnit(key) {
  var quantityUnit =  AgrochemicalQuantityUnits.find(item => item.key === key.toLowerCase())
  return quantityUnit ? quantityUnit.label : ''
}

export const ChemicalTypes = [
  { key: 'DISINFECTANT',  label: 'Дезинфицирующее средство' },
  { key: 'FERTILIZER', label: 'Удобрение' },
  { key: 'HORMONE', label: 'Гормон и активатор роста' },
  { key: 'MANURE', label: 'Навоз' },
  { key: 'PESTICIDE', label: 'Пестицид' },
]

export function FindChemicalType(key) {
  var chemicalType =  ChemicalTypes.find(item => item.key === key.toLowerCase())
  return chemicalType ? chemicalType.label : ''
}

export const GrowingMediumQuantityUnits = [
  { key: 'BAGS',  label: 'Мешки' },
  { key: 'CUBIC_METRE', label: 'Кубический метр' },
]

export function FindGrowingMediumQuantityUnit(key) {
  var quantityUnit =  GrowingMediumQuantityUnits.find(item => item.key === key.toLowerCase())
  return quantityUnit ? quantityUnit.label : ''
}


export const PlantUnits = [
  { key: 'UNITS',  label: 'Units' },
  { key: 'PACKETS', label: 'Packets' },
]

export function FindPlantUnit(key) {
  var quantityUnit =  PlantUnits.find(item => item.key === key.toLowerCase())
  return quantityUnit ? quantityUnit.label : ''
}
