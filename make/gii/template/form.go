/**
 * @Time : 2020/10/9 17:47
 * @Author : liang
 * @File : form
 * @Software: GoLand
 */

package template

const Form = `<?php


namespace %s;

use common\librarys\BaseForm;

class %sForm extends BaseForm 
{

    /**
     * @inheritDoc
     */
    protected function rules(): array
    {
        return [

        ];
    }
}
`
