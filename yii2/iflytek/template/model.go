/**
 * @Time : 2020/10/9 17:12
 * @Author : liang
 * @File : model
 * @Software: GoLand
 */

package template

const Model = `<?php


namespace %s;


use Yii;
use yii2\db\ActiveRecord;

%s
class %s extends ActiveRecord
{
    /**
     * @inheritDoc
     */
    public static function getDb()
    {
        return Yii::$app->db_v2;
    }

    /**
     * @return string
     */
    public static function tableName()
    {
        return '%s';
    }

    /**
     * @return array
     */
    public function rules()
    {
        return [

        ];
    }


}

`
