package api

import (
	"fmt"
	"game/config"
	"log"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Attack(c *gin.Context) {

	/* mock data
	{
		"AttackCity":"1",
		"AttackedCity":"2"
	}
	*/

	type commandData struct {
		AttackCity   string `json:"AttackCity"`
		AttackedCity string `json:"AttackedCity"`
	}

	type armOfCity struct {
		Id                 string
		Name               string  //城市名
		Army               float64 //兵力
		AverageCombatPower float64 //平均战力
		Power              float64 //总战斗力
		Remain             int     //剩余兵力
	}
	var i int = 0
	var cmd commandData           //将用户传入的攻击城市和被攻击城市绑定到cmd参数
	var cityInfo = [2]armOfCity{} //定义一个不定长度的空切片

	//!!!!!!!!!!!!!想办法把commadData结构杀掉，全都用armofcity多优雅
	if err := c.BindJSON(&cmd); err != nil {
		log.Fatal(err)
	}

	if cmd.AttackCity == cmd.AttackedCity {
		c.IndentedJSON(200, gin.H{
			"msg": "攻击者和被攻击者不能为同一城市",
		})
	}

	cityInfo[0].Id = cmd.AttackCity
	cityInfo[1].Id = cmd.AttackedCity

	//获取两座城市的基本信息
	conn := config.Dbconnect()
	rows, err := conn.Query("select name,army,averageCombatPower from city where id=? or id=?", cmd.AttackCity, cmd.AttackedCity)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&cityInfo[i].Name, &cityInfo[i].Army, &cityInfo[i].AverageCombatPower); err != nil {
			log.Fatal(err)
		}
		//***重要算法***计算两军战斗力
		cityInfo[i].Power = math.Pow(cityInfo[i].Army, 2) * cityInfo[i].AverageCombatPower
		i++
	}
	fmt.Println(cityInfo[0])
	fmt.Println(cityInfo[1])

	//计算两军剩余人数 *****重要算法*****
	cityInfo[0].Remain = int(math.Sqrt(cityInfo[0].Power-0.1*cityInfo[1].Power) / cityInfo[1].AverageCombatPower)
	cityInfo[1].Remain = int(math.Sqrt(cityInfo[1].Power-0.1*cityInfo[0].Power) / cityInfo[0].AverageCombatPower)

	// 更新city表army列
	for _, i := range cityInfo {
		if _, err = conn.Query("update city set army=? where id=?", i.Remain, i.Id); err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(200, gin.H{
		"remain0":         cityInfo[0].Remain,
		"remain1":         cityInfo[1].Remain,
		"lostPercentage0": strconv.Itoa(100*(int(cityInfo[0].Army)-cityInfo[0].Remain)/int(cityInfo[0].Army)) + "%",
		"lostPercentage1": strconv.Itoa(100*(int(cityInfo[1].Army)-cityInfo[1].Remain)/int(cityInfo[1].Army)) + "%",
	})
}
