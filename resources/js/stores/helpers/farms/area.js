export const AreaTypes = [
  { key: 'SEEDING', label: 'Посев' },
  { key: 'GROWING', label: 'Выращивание' },
];

export const AreaLocations = [
  { key: 'OUTDOOR', label: 'Поле (Outdoor)' },
  { key: 'INDOOR', label: 'Теплица (Indoor)' },
];

export const AreaSizeUnits = [
  { key: 'Ha', label: 'Га' },
  { key: 'm2', label: 'м2' },
];

export function FindAreaType(key) {
  return AreaTypes.find(item => item.key === key);
}

export function FindAreaLocation(key) {
  return AreaLocations.find(item => item.key === key);
}

export function FindAreaSizeUnit(key) {
  return AreaSizeUnits.find(item => item.key === key);
}
