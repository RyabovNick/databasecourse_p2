'use strict'

module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.bulkInsert(
      'types',
      [
        {
          name: 'Членистоногие',
          description:
            'Членистоно́гие (лат. Arthropoda, от др.-греч. ἄρθρον — сустав и πούς, род. п. ποδός — нога) — тип первичноротых животных, включающий насекомых, ракообразных, паукообразных и многоножек. По количеству видов и распространённости может считаться самой процветающей группой живых организмов. К представителям этого типа относится 2/3 всех видов живых существ на Земле.',
          created_at: new Date(),
          updated_at: new Date(),
        },
        {
          name: 'Хордовые',
          description:
            'Хо́рдовые (лат. Chordata) — тип вторичноротых животных, для которых характерно наличие энтодермального осевого скелета в виде хорды, которая у высших форм заменяется позвоночником. По строению и функции нервной системы тип хордовых занимает высшее место среди всех животных. В мире известно более 60 000 видов хордовых.',
          created_at: new Date(),
          updated_at: new Date(),
        },
        {
          name: 'Волосатики',
          description:
            'Волоса́тики[1] (лат. Nematomorpha, от др.-греч. νῆμα, родительный падеж νῆματος — нить, μορφή — форма) — тип беспозвоночных животных, личинки которых ведут паразитоидный образ жизни. В ископаемом виде известны с эоцена.',
          created_at: new Date(),
          updated_at: new Date(),
        },
      ],
      {},
    )
  },

  down: (queryInterface, Sequelize) => {
    return queryInterface.bulkDelete('types', null, {})
    /*
      Add reverting commands here.
      Return a promise to correctly handle asynchronicity.

      Example:
      return queryInterface.bulkDelete('People', null, {});
    */
  },
}
