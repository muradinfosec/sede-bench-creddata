// export const table: {
//   epoch: number;
//   secret: string;
//   token: string;
// }[] = [
//   {
//     epoch: 1565103854545,
//     secret: 'PBOU2OAHI5RM5M6PJU6CF7DJQXFZM36E',
//     token: '566155'
//   },
//   {
//     epoch: 1565103903110,
//     secret: 'SOMNOXCAGK7UG4TUA4TMSRXLELWWSW3C',
//     token: '540849'
//   },
//   {
//     epoch: 1565106003151,
//     secret: 'HR4K6BGZBTAY64HVDWSEK7BSOCSMJESY',
//     token: '668733'
//   },
//   {
//     epoch: 1565106018408,
//     secret: 'VLEKIDDWFA5WU5YSHYVBK6WOCYXOH2HW',
//     token: '767234'
//   },
//   {
//     epoch: 1565106407848,
//     secret: 'SN5AP2XBM56RQLREFDQYW5ACSBHAJ4WM',
//     token: '942732'
//   },
//   {
//     epoch: 1565106431089,
//     secret: 'QG6R7HESZYQO2R2TCOUHLQHYMN2Z4WJA',
//     token: '235413'
//   },
//   {
//     epoch: 1565106483557,
//     secret: 'K4ZSZVKWH6AHTASEVBZWFU6LQYOGDCRJ',
//     token: '508543'
//   }
// ];

export const deltaTable = {
  secret: 'R23FLMIHSPXL67WAOH3MWKNZIW2FOELD',
  epoch: 1565973031233,
  rows: [
    ['039223', -2],
    ['311336', -1],
    ['288367', 0],
    ['408608', 1],
    ['721767', 2],
    ['819412', null]
  ] as [string, number | null][]
};

export const table: {
  decoded: string;
  digest: string;
  secret: string;
  epoch: number;
  token: string;
}[] = [
  {
    decoded: '68442f372b67474e2f47617679706f6e30756f51',
    digest: '422eb1a849cf0650fef4dbdd8b0ee0fe57a87eb9',
    epoch: 1565103854545,
    secret: 'BFWS6RYQN2NB5H5ZTM5WU6PZCQAVT32B',
    token: '566155'
  },
  {
    decoded: '68442f372b67474e2f47617679706f6e30756f51',
    digest: 'c305b82dbf2a8d2d8a22e9d3992e4e666222d0e2',
    secret: 'CUAH4JXSB3NI7O5EJA2DQ2NDGBYLH43Z',
    epoch: 1565103878581,
    token: '522154'
  },
  {
    decoded: '636c6c4e506479436f314f6b4852623167564f76',
    digest: '64a959e511420af1a406424f87b4412977b3cbd4',
    secret: 'BCSCHQVAFG4JD6CBV7IXGBXTOMJWKP3W',
    epoch: 1565103903110,
    token: '540849'
  }
];
